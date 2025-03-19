use crate::native::types::JsInputs;
use std::collections::HashMap;

#[napi(object)]
/// Stripped version of the KhulnasoftJson interface for use in rust
pub struct KhulnasoftJson {
    pub named_inputs: Option<HashMap<String, Vec<JsInputs>>>,
}
