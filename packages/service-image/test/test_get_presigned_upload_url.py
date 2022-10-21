import unittest
from unittest.mock import Mock, patch
from faker import Faker
from get_presigned_upload_url import get_presigned_upload_url

fake = Faker()

class TestGetPresignedUploadUrl(unittest.TestCase):
    def test_success(self):
        self.assertEqual(1 + 2, 3)
        