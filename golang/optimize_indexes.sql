-- ISUCONパフォーマンス最適化用インデックス
-- このファイルをMySQLで実行してください

-- posts テーブルの最適化
ALTER TABLE posts ADD INDEX idx_created_at (created_at DESC);
ALTER TABLE posts ADD INDEX idx_user_id_created_at (user_id, created_at DESC);

-- comments テーブルの最適化
ALTER TABLE comments ADD INDEX idx_post_id_created_at (post_id, created_at DESC);
ALTER TABLE comments ADD INDEX idx_user_id (user_id);

-- users テーブルの最適化
ALTER TABLE users ADD INDEX idx_account_name_del_flg (account_name, del_flg);

-- 複合インデックスで更なる最適化
ALTER TABLE posts ADD INDEX idx_id_user_id (id, user_id); 
