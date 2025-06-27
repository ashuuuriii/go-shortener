create table shortened_urls (
    created_at timestamp with time zone not null default now(),
    short_url text unique not null,
    long_url text not null,
    expire_in_days integer not null
);

create index idx_shortened_urls_shortened_url on shortened_urls using hash (short_url);

grant all privileges on public.shortened_urls to "shortener-user";