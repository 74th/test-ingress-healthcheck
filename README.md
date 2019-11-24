# GKE Ingressのヘルスチェックが自動でパスが変わるか

概要

- Ingressを作ると、ロードバランサ、バックエンドサービスが作られる。バックエンドサービスにヘルスチェックの設定がある。
- Ingressにはヘルスチェックのパスを、PodのReadinessProbeの値を参照する、という機能がある。これがどのように動作するか確認する。

わかったこと

- readinessProbeのGETのパスが使われる。
- 複数のコンテナがある場合、serviceで指定したポートを出力するコンテナのreadinessProbeが使われる。
- Cloud Endpointsでは、そのヘルスチェックのAPIもopenapiに追加する必要がある。
- Cloud Endpointsのコンテナが仲介する場合、そのESPにreadinessProbeを設定する。
