output "registry-id" {
  value = yandex_iot_core_registry.iot-registry.id
}
output "device-id" {
  value = yandex_iot_core_device.rpi-device.id
}

output "device-password" {
  value = random_password.rpi-device-password.result
}

output "registry-password" {
  value = random_password.iot-registry-password.result
}