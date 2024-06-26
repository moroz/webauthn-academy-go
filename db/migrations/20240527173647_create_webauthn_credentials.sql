-- +goose Up
-- +goose StatementBegin
create table webauthn_credentials (
  id bigint primary key generated by default as identity,
  webauthn_id bytea not null unique,
  user_id bigint not null references users (id) on delete cascade,
  public_key bytea not null,
  attestation_type text,
  transport text[],
  user_present bool not null default false,
  user_verified bool not null default false,
  backup_eligible bool not null default false,
  backup_state bool not null default false,
  authenticator_aaguid bytea,
  authenticator_sign_count int not null default 0,
  authenticator_clone_warning boolean not null default false,
  authenticator_attachment text,
  inserted_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table webauthn_credentials;
-- +goose StatementEnd
