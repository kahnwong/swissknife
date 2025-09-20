use battery;
use log::error;

#[repr(C)]
#[derive(Debug)]
pub enum BatteryTimeToEmptyError {
    Success = 0,
    NoBattery = 1,
    NoTimeToEmpty = 2,
    ManagerError = 3,
}

#[repr(C)]
pub struct BatteryTimeToEmptyResult {
    pub time_to_empty_seconds: u64,
    pub error: BatteryTimeToEmptyError,
}

#[no_mangle]
pub extern "C" fn battery_time_to_empty() -> BatteryTimeToEmptyResult {
    match battery::Manager::new() {
        Ok(manager) => {
            match manager.batteries() {
                Ok(batteries) => {
                    // Get the first battery if available
                    if let Some(battery) = batteries.into_iter().next() {
                        match battery {
                            Ok(bat) => {
                                match bat.time_to_empty() {
                                    Some(time) => {
                                        // Convert Time to seconds as u64
                                        let seconds = time.get::<battery::units::time::second>();
                                        BatteryTimeToEmptyResult {
                                            time_to_empty_seconds: seconds as u64,
                                            error: BatteryTimeToEmptyError::Success,
                                        }
                                    }
                                    None => {
                                        error!("No time to empty available");
                                        BatteryTimeToEmptyResult {
                                            time_to_empty_seconds: 0,
                                            error: BatteryTimeToEmptyError::NoTimeToEmpty,
                                        }
                                    }
                                }
                            }
                            Err(e) => {
                                error!("Battery error: {}", e);
                                BatteryTimeToEmptyResult {
                                    time_to_empty_seconds: 0,
                                    error: BatteryTimeToEmptyError::NoBattery,
                                }
                            }
                        }
                    } else {
                        error!("No batteries found");
                        BatteryTimeToEmptyResult {
                            time_to_empty_seconds: 0,
                            error: BatteryTimeToEmptyError::NoBattery,
                        }
                    }
                }
                Err(e) => {
                    error!("Failed to get batteries: {}", e);
                    BatteryTimeToEmptyResult {
                        time_to_empty_seconds: 0,
                        error: BatteryTimeToEmptyError::ManagerError,
                    }
                }
            }
        }
        Err(e) => {
            error!("Failed to create battery manager: {}", e);
            BatteryTimeToEmptyResult {
                time_to_empty_seconds: 0,
                error: BatteryTimeToEmptyError::ManagerError,
            }
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_battery_time_to_empty() {
        let result = battery_time_to_empty();

        match result.error {
            BatteryTimeToEmptyError::Success => {
                println!(
                    "Battery time to empty: {} seconds",
                    result.time_to_empty_seconds
                );
                assert!(result.time_to_empty_seconds > 0);
            }
            BatteryTimeToEmptyError::NoBattery => {
                println!("No battery found");
                assert_eq!(result.time_to_empty_seconds, 0);
            }
            BatteryTimeToEmptyError::NoTimeToEmpty => {
                println!("No time to empty available");
                assert_eq!(result.time_to_empty_seconds, 0);
            }
            BatteryTimeToEmptyError::ManagerError => {
                println!("Battery manager error");
                assert_eq!(result.time_to_empty_seconds, 0);
            }
        }

        assert!(matches!(
            result.error,
            BatteryTimeToEmptyError::Success
                | BatteryTimeToEmptyError::NoBattery
                | BatteryTimeToEmptyError::NoTimeToEmpty
                | BatteryTimeToEmptyError::ManagerError
        ));
    }
}
