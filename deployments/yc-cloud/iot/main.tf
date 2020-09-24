
resource "random_password" "rpi-device-password" {
  length = 16
}
resource "random_password" "iot-registry-password" {
  length = 16
}
resource "yandex_iot_core_registry" "iot-registry" {
  name = var.iot-registry-name
  passwords = [random_password.iot-registry-password.result]
}
resource "yandex_iot_core_device" "rpi-device" {
  name = var.iot-device-name
  registry_id = yandex_iot_core_registry.iot-registry.id
  passwords = [random_password.rpi-device-password.result]
}