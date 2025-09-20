extern crate battery;
extern crate log;
extern crate sysinfo;

mod battery_cycle_count;
mod battery_time_to_empty;
mod sensors;

pub use battery_cycle_count::{battery_cycle_count, BatteryError, BatteryResult};
pub use battery_time_to_empty::{
    battery_time_to_empty, BatteryTimeToEmptyError, BatteryTimeToEmptyResult,
};
pub use sensors::{sensors, SensorError, SensorResult};
