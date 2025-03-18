import { Hooks } from "./types.js";
import {KhulnasoftCustomHook} from "./khulnasoft-custom-hook.js";

/*
 * This file is only ever generated once on the first generation and then is free to be modified.
 * Any hooks you wish to add should be registered in the initHooks function. Feel free to define them
 * in this file or in separate files in the hooks folder.
 */

export function initHooks(hooks: Hooks) {
    hooks.registerAfterSuccessHook(new KhulnasoftCustomHook())
    hooks.registerBeforeRequestHook(new KhulnasoftCustomHook())
    hooks.registerBeforeCreateRequestHook(new KhulnasoftCustomHook())
}
