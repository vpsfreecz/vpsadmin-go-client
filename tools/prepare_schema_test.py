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
                "resources": {
                    "example": {"default": "unchanged"},
                    "vps": {"resources": {"evidence": self.evidence_resource(
                        "/v7.0/vpses/{vps_id}/evidence", "vps.evidence#index", "evidence"
                    )}},
                    "vps_evidence": self.evidence_resource(
                        "/v7.0/vps_evidence", "vps_evidence#index", "vps_evidence"
                    ),
                },
            },
        }

    @staticmethod
    def evidence_resource(path, scope, namespace):
        return {"resources": {}, "actions": {"index": {
            "method": "GET", "path": path, "scope": scope, "auth": True,
            "input": {"layout": "object", "namespace": namespace},
            "output": {"layout": "object_list", "namespace": namespace},
        }}}

    def prepare(self, schema):
        self.source.write_text(json.dumps(schema), encoding="utf-8")
        return subprocess.run(
            [sys.executable, str(HELPER), str(self.source), str(self.output)],
            capture_output=True, text=True,
        )

    def test_only_two_projections_change_and_source_bytes_are_preserved(self):
        result = self.prepare(self.schema)
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(self.source.read_text(), json.dumps(self.schema))
        expected = copy.deepcopy(self.schema)
        expected["description"]["authentication"]["oauth2"]["revoke_url"] = (
            "https://auth.vpsfree.cz/_auth/oauth2/revoke"
        )
        del expected["description"]["resources"]["vps"]["resources"]["evidence"]
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

    def test_changed_evidence_signatures_preserve_existing_output(self):
        for key in ("child", "fleet"):
            for field, value in (
                ("method", "POST"), ("path", "/wrong"), ("scope", "wrong#index"),
                ("auth", False), ("auth", 1),
                ("input.layout", "hash"), ("input.namespace", "wrong"),
                ("output.layout", "object"), ("output.namespace", "wrong"),
            ):
                with self.subTest(resource=key, field=field, value=value):
                    schema = copy.deepcopy(self.schema)
                    resources = schema["description"]["resources"]
                    resource = (resources["vps"]["resources"]["evidence"]
                                if key == "child" else resources["vps_evidence"])
                    target = resource["actions"]["index"]
                    parts = field.split(".")
                    if len(parts) == 2:
                        target = target[parts[0]]
                    target[parts[-1]] = value
                    self.output.write_text("existing output")
                    self.assertNotEqual(self.prepare(schema).returncode, 0)
                    self.assertEqual(self.output.read_text(), "existing output")

            for change in ("missing", "extra_action", "extra_child"):
                with self.subTest(resource=key, change=change):
                    schema = copy.deepcopy(self.schema)
                    resources = schema["description"]["resources"]
                    parent = resources["vps"]["resources"] if key == "child" else resources
                    name = "evidence" if key == "child" else "vps_evidence"
                    if change == "missing":
                        del parent[name]
                    elif change == "extra_action":
                        parent[name]["actions"]["show"] = {}
                    else:
                        parent[name]["resources"]["unexpected"] = {}
                    self.output.write_text("existing output")
                    self.assertNotEqual(self.prepare(schema).returncode, 0)
                    self.assertEqual(self.output.read_text(), "existing output")


if __name__ == "__main__":
    unittest.main()
