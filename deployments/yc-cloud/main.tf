
provider "yandex" {

}
module "iot" {
  source = "./iot"
  iot-device-name = var.iot-device-name
  iot-registry-name = var.iot-registry-name
}

module "function" {
  source = "./function"
  yc_folder = var.yc-folder
  yc_cloud = var.yc-cloud
}