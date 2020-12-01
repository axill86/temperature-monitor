resource "yandex_iam_service_account" "function-account" {
  name = "function-account"
  description = "Service account to run a key"
}

resource "yandex_resourcemanager_folder_iam_binding" "service-account-binding" {
  folder_id = var.yc_folder
  members = [
    "serviceAccount:${yandex_iam_service_account.function-account.id}"
  ]
  role = "serverless.functions.invoker"
}