use obi::{OBIDecode, OBIEncode, OBISchema};
use owasm_kit::{execute_entry_point, ext, oei, prepare_entry_point};

#[derive(OBIDecode, OBISchema)]
struct Input {
    rpc: String,
    to: String,
}

#[derive(OBIEncode, OBISchema)]
struct Output {
    total_supply: String,
}

const DATA_SOURCE_ID: i64 = 1;
const EXTERNAL_ID: i64 = 0;

#[no_mangle]
fn prepare_impl(input: Input) {
    oei::ask_external_data(
        EXTERNAL_ID,
        DATA_SOURCE_ID,
        format!("{} {}", input.rpc, input.to).as_bytes(),
    );
}

#[no_mangle]
fn execute_impl(_input: Input) -> Output {
    Output {
        total_supply: ext::load_majority::<String>(EXTERNAL_ID).unwrap(),
    }
}

prepare_entry_point!(prepare_impl);
execute_entry_point!(execute_impl);
