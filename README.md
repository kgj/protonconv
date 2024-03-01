# protonconv
Tool to convert proton pass exports to keepass db files

This is not guaranteed to work in any capacity - it worked for me and my protonpass export.

Uses github.com/tobischo/gokeepasslib for keepass integration

Provides CLI and GUI user interfaces - if run with no flags the GUI version is executed.

There are 3 parameters that can be passed via the commandline:

in - This is the full path to the json export 

db - This is the full path to the destination keepass

pass - This is the new master password for your new keepass db

Note that if you use an existing keepass db this will silently overwrite it.
