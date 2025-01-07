\set pguser `echo "$POSTGRES_USER"`

CREATE DATABASE _khulnasoft WITH OWNER :pguser;
