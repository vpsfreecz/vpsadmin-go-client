import copy
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest


HELPER = Path(__file__).with_name("prepare-schema")


class PrepareSchemaTest(unittest.TestCase):
    def setUp(self):
        self.directory = tempfile.TemporaryDirectory()
        self.addCleanup(self.directory.cleanup)
        root = Path(self.directory.name)
        self.source = root / "source.json"
        self.output = root / "effective.json"
        self.schema = {
            "schema_format": 1,
            "api_version": "7.0",
            "protocol_version": "2.0",
            "profile": "vpsadmin-full-fixture",
            "mount_path": "/",
            "source_url": "https://schema.invalid/v7.0/",
            "description": {
                "authentication": {
                    "oauth2": {"revoke_url": "http://api.test/_auth/oauth2/revoke"}
                },
                "resources": {"example": {"default": "unchanged"}},
            },
        }

    def prepare(self, schema):
        self.source.write_text(json.dumps(schema), encoding="utf-8")
        return subprocess.run(
            [sys.executable, str(HELPER), str(self.source), str(self.output)],
            capture_output=True, text=True,
        )

    def test_only_endpoint_changes_and_source_bytes_are_preserved(self):
        result = self.prepare(self.schema)
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(self.source.read_text(), json.dumps(self.schema))
        expected = copy.deepcopy(self.schema)
        expected["description"]["authentication"]["oauth2"]["revoke_url"] = (
            "https://auth.vpsfree.cz/_auth/oauth2/revoke"
        )
        self.assertEqual(json.loads(self.output.read_text()), expected)

    def test_repeat_is_byte_identical(self):
        self.assertEqual(self.prepare(self.schema).returncode, 0)
        first = self.output.read_bytes()
        self.assertEqual(self.prepare(self.schema).returncode, 0)
        self.assertEqual(self.output.read_bytes(), first)

    def test_unexpected_inputs_preserve_existing_output(self):
        for field, value in (
            ("schema_format", True), ("schema_format", 2),
            ("api_version", "8.0"), ("protocol_version", "2.1"),
            ("profile", "public"), ("mount_path", "/other/"),
            ("source_url", "https://example.test/v7.0/"),
        ):
            with self.subTest(field=field, value=value):
                schema = copy.deepcopy(self.schema)
                schema[field] = value
                self.output.write_text("existing output")
                result = self.prepare(schema)
                self.assertNotEqual(result.returncode, 0)
                self.assertEqual(self.output.read_text(), "existing output")

        for value in ("https://other.test/revoke", None):
            with self.subTest(revoke_url=value):
                schema = copy.deepcopy(self.schema)
                schema["description"]["authentication"]["oauth2"]["revoke_url"] = value
                self.assertNotEqual(self.prepare(schema).returncode, 0)
                self.assertEqual(self.output.read_text(), "existing output")

    def test_source_cannot_be_overwritten(self):
        original = json.dumps(self.schema)
        self.source.write_text(original)
        result = subprocess.run(
            [sys.executable, str(HELPER), str(self.source), str(self.source)],
            capture_output=True, text=True,
        )
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(self.source.read_text(), original)


if __name__ == "__main__":
    unittest.main()
