
provider "yandex" {

}
module "iot" {
  source = "./iot"
  iot-device-name = var.iot-device-name
  iot-registry-name = var.iot-registry-name
}