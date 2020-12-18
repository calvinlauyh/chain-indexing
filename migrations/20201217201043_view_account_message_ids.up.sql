CREATE TABLE view_account_message_ids (
    account VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    message_index INT NOT NULL,
    PRIMARY KEY(account, transaction_hash, message_index)
)