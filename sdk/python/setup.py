# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call

class InstallPluginCommand(install):
    def run(self):
        install.run(self)
        check_call(['pulumi', 'plugin', 'install', 'resource', 'openstack', '${PLUGIN_VERSION}'])

setup(name='pulumi_openstack',
      version='${VERSION}',
      description='A Pulumi package for creating and managing OpenStack cloud resources.',
      cmdclass={
          'install': InstallPluginCommand,
      },
      keywords='pulumi openstack',
      url='https://pulumi.io',
      project_urls={
          'Repository': 'https://github.com/pulumi/pulumi-openstack'
      },
      license='Apache 2.0',
      packages=find_packages(),
      install_requires=[
          'pulumi>=0.14.0,<0.15.0'
      ],
      zip_safe=False)
