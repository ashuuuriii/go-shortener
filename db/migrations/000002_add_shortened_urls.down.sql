revoke all privileges on table public.shortened_urls from "shortener-user";

drop index idx_shortened_urls_shortened_url;

drop table shortened_urls;