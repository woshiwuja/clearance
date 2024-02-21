CREATE TABLE devices (
id text PRIMARY KEY NOT NULL,
name text NOT NULL,
model text NOT NULL,
ip_addr text NOT NULL,
mac_addr text NOT NULL
);

CREATE TABLE events (
id UUID PRIMARY KEY,
description text,
source_addr text,
dst_addr text,
category text,
event_code text,
is_offence bool --Is offence if matches rule
);

CREATE TABLE offences (
id UUID PRIMARY KEY,
description text,
source_addr text,
dst_addr text,
category text,
event_code text
);

CREATE TABLE rules(
id UUID PRIMARY KEY,
severity int,
relevance int,
credibility int,
is_logged bool,
category text,
of_group text
);

CREATE TABLE categories (
id UUID PRIMARY KEY,
name text);
