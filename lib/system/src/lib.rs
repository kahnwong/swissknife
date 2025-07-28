extern crate log;
extern crate sysinfo;

mod sensors;
pub use sensors::{sensors, SensorError, SensorResult};
