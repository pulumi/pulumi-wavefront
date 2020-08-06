"""A Python Pulumi program"""

import pulumi
import pulumi_wavefront as wavefront

user = wavefront.User("demo-py",
                      email="test+py@mycompany.io")
