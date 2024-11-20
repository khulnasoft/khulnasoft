import khulnasoft
import khulnasoft_random as random

random_password = random.RandomPassword("randomPassword",
    length=16,
    special=True,
    override_special="_%@")
khulnasoft.export("password", random_password.result)
