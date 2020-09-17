output "registry-id" {
  value = yandex_iot_core_registry.iot-registry.id
}
output "device-id" {
  value = yandex_iot_core_device.rpi-device.id
}