provider "yandex" {

}
resource "yandex_iot_core_registry" "iot-registry" {
  name = var.iot-registry-name
}
resource "yandex_iot_core_device" "rpi-device" {
  name = var.iot-device-name
  registry_id = yandex_iot_core_registry.iot-registry.id
}