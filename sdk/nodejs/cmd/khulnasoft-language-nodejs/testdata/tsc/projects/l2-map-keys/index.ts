import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as plain from "@khulnasoft/plain";
import * as primitive from "@khulnasoft/primitive";
import * as primitive_ref from "@khulnasoft/primitive-ref";
import * as ref_ref from "@khulnasoft/ref-ref";

const prim = new primitive.Resource("prim", {
    boolean: false,
    float: 2.17,
    integer: -12,
    string: "Goodbye",
    numberArray: [
        0,
        1,
    ],
    booleanMap: {
        "my key": false,
        "my.key": true,
        "my-key": false,
        my_key: true,
        MY_KEY: false,
        myKey: true,
    },
});
const ref = new primitive_ref.Resource("ref", {data: {
    boolean: false,
    float: 2.17,
    integer: -12,
    string: "Goodbye",
    boolArray: [
        false,
        true,
    ],
    stringMap: {
        "my key": "one",
        "my.key": "two",
        "my-key": "three",
        my_key: "four",
        MY_KEY: "five",
        myKey: "six",
    },
}});
const rref = new ref_ref.Resource("rref", {data: {
    innerData: {
        boolean: false,
        float: -2.17,
        integer: 123,
        string: "Goodbye",
        boolArray: [],
        stringMap: {
            "my key": "one",
            "my.key": "two",
            "my-key": "three",
            my_key: "four",
            MY_KEY: "five",
            myKey: "six",
        },
    },
    boolean: true,
    float: 4.5,
    integer: 1024,
    string: "Hello",
    boolArray: [],
    stringMap: {
        "my key": "one",
        "my.key": "two",
        "my-key": "three",
        my_key: "four",
        MY_KEY: "five",
        myKey: "six",
    },
}});
const plains = new plain.Resource("plains", {
    data: {
        innerData: {
            boolean: false,
            float: 2.17,
            integer: -12,
            string: "Goodbye",
            boolArray: [
                false,
                true,
            ],
            stringMap: {
                "my key": "one",
                "my.key": "two",
                "my-key": "three",
                my_key: "four",
                MY_KEY: "five",
                myKey: "six",
            },
        },
        boolean: true,
        float: 4.5,
        integer: 1024,
        string: "Hello",
        boolArray: [
            true,
            false,
        ],
        stringMap: {
            "my key": "one",
            "my.key": "two",
            "my-key": "three",
            my_key: "four",
            MY_KEY: "five",
            myKey: "six",
        },
    },
    nonPlainData: {
        innerData: {
            boolean: false,
            float: 2.17,
            integer: -12,
            string: "Goodbye",
            boolArray: [
                false,
                true,
            ],
            stringMap: {
                "my key": "one",
                "my.key": "two",
                "my-key": "three",
                my_key: "four",
                MY_KEY: "five",
                myKey: "six",
            },
        },
        boolean: true,
        float: 4.5,
        integer: 1024,
        string: "Hello",
        boolArray: [
            true,
            false,
        ],
        stringMap: {
            "my key": "one",
            "my.key": "two",
            "my-key": "three",
            my_key: "four",
            MY_KEY: "five",
            myKey: "six",
        },
    },
});
