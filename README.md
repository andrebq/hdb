# What is this?

**HDB** exposes an API similar to a filesystem using a CAS storage to hold
all the data.

After data is inserted into **HDB** it is immutable. To make **updates** to
a file or directory the user will create a new blob in the database and update
the **current** reference to the new **ref**.