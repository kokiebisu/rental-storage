import unittest
from faker import Faker

fake = Faker()

class TestAdd(unittest.TestCase):
    def test_success(self):
        self.assertEqual(1 + 2, 3)
        