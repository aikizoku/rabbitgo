package util

// 400 Bad Request タイプミス等、リクエストにエラーがあります。
// 401 Unauthorixed  認証に失敗しました。（パスワードを適当に入れてみた時などに発生）
// 402 Payment Required  （将来の使用のための予約コード）
// 403 Forbidden あなたにはアクセス権がありません。
// 404 (File)Not Found 該当アドレスのページはありません、またはそのサーバーが落ちている。
// 405 Method Not Allowed  許可されていないメソッドタイプのリクエストを受けた。
// 406 Not Acceptable  Acceptヘッダから判断された結果、受け入れられない内容を持っていた。
// 407 Proxy Authentication Required 最初にProxy認証が必要です。
// 408 Request Time-out  リクエストの待ち時間内に反応がなかった。
// 409 Conflict  そのリクエストは現在の状態のリソースと矛盾するため完了できなかった。
// 410 Gone  そのリクエストはサーバでは利用できず転送先のアドレスも分からない。
// 411 Length Required 定義されたContent-Lengthの無いリクエストを拒否しました。
// 412 Precondition Failed 1つ以上のリクエストヘッダフィールドで与えられた条件がサーバ上のテストで不正であると判断しました。
// 413 Request Entity Too Large  処理可能量より大きいリクエストのため拒否しました。
// 414 Request-URI Too Large リクエストURIが長すぎるため拒否しました。
// 415 Unsupported Media Type  リクエストされたメソッドに対してリクエストされたリソースがサポートしていないフォーマットであるため、サーバはリクエストのサービスを拒否しました。
// 416 Requested range not satisfiable リクエストにRangeヘッダフィールドは含まれていたが、If-Rangeリクエストヘッダフィールドがありません。
// 417 Expectation Failed  Expectリクエストヘッダフィールド拡張が受け入れられなかった。
