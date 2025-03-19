use std::sync::{Condvar, LockResult, Mutex, MutexGuard};

pub struct KhulnasoftMutex<T>(Mutex<T>);

impl<T> KhulnasoftMutex<T> {
    pub fn new(value: T) -> Self {
        Self(Mutex::new(value))
    }
    pub fn lock(&self) -> LockResult<MutexGuard<'_, T>> {
        self.0.lock()
    }
}

pub struct KhulnasoftCondvar(Condvar);

impl KhulnasoftCondvar {
    pub fn new() -> Self {
        Self(Condvar::new())
    }

    pub fn wait<'a, T, F>(
        &self,
        mutex_guard: MutexGuard<'a, T>,
        condition: F,
    ) -> LockResult<MutexGuard<'a, T>>
    where
        F: Fn(&mut T) -> bool,
    {
        self.0.wait_while(mutex_guard, condition)
    }
}
