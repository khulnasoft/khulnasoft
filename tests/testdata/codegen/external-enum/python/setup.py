# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "0.0.0"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "example Pulumi Package - Development Version"


setup(name='khulnasoft_example',
      python_requires='>=3.8',
      version=VERSION,
      long_description=readme(),
      long_description_content_type='text/markdown',
      packages=find_packages(),
      package_data={
          'khulnasoft_example': [
              'py.typed',
              'khulnasoft-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'khulnasoft>=3.0.0,<4.0.0',
          'khulnasoft-google-native>=0.20.0,<1.0.0',
          'semver>=2.8.1',
          'typing-extensions>=4.11,<5; python_version < "3.11"'
      ],
      zip_safe=False)