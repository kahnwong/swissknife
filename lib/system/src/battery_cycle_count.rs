use battery;
use log::error;

#[repr(C)]
#[derive(Debug)]
pub enum BatteryError {
    Success = 0,
    NoBattery = 1,
    NoCycleCount = 2,
    ManagerError = 3,
}

#[repr(C)]
pub struct BatteryResult {
    pub cycle_count: u32,
    pub error: BatteryError,
}

#[no_mangle]
pub extern "C" fn battery_cycle_count() -> BatteryResult {
    match battery::Manager::new() {
        Ok(manager) => {
            match manager.batteries() {
                Ok(batteries) => {
                    // Get the first battery if available
                    if let Some(battery) = batteries.into_iter().next() {
                        match battery {
                            Ok(bat) => {
                                match bat.cycle_count() {
                                    Some(count) => BatteryResult {
                                        cycle_count: count,
                                        error: BatteryError::Success,
                                    },
                                    None => {
                                        error!("No cycle count available");
                                        BatteryResult {
                                            cycle_count: 0,
                                            error: BatteryError::NoCycleCount,
                                        }
                                    }
                                }
                            }
                            Err(e) => {
                                error!("Battery error: {}", e);
                                BatteryResult {
                                    cycle_count: 0,
                                    error: BatteryError::NoBattery,
                                }
                            }
                        }
                    } else {
                        error!("No batteries found");
                        BatteryResult {
                            cycle_count: 0,
                            error: BatteryError::NoBattery,
                        }
                    }
                }
                Err(e) => {
                    error!("Failed to get batteries: {}", e);
                    BatteryResult {
                        cycle_count: 0,
                        error: BatteryError::ManagerError,
                    }
                }
            }
        }
        Err(e) => {
            error!("Failed to create battery manager: {}", e);
            BatteryResult {
                cycle_count: 0,
                error: BatteryError::ManagerError,
            }
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_battery_cycle_count() {
        let result = battery_cycle_count();

        match result.error {
            BatteryError::Success => {
                println!("Battery cycle count: {}", result.cycle_count);
                assert!(result.cycle_count >= 1);
            }
            BatteryError::NoBattery => {
                println!("No battery found");
                assert_eq!(result.cycle_count, 0);
            }
            BatteryError::NoCycleCount => {
                println!("No cycle count available");
                assert_eq!(result.cycle_count, 0);
            }
            BatteryError::ManagerError => {
                println!("Battery manager error");
                assert_eq!(result.cycle_count, 0);
            }
        }

        assert!(matches!(
            result.error,
            BatteryError::Success
                | BatteryError::NoBattery
                | BatteryError::NoCycleCount
                | BatteryError::ManagerError
        ));
    }
}
