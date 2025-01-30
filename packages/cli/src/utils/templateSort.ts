import * as sdk from 'khulnasoft'

export function sortTemplatesAliases<
  E extends sdk.components['schemas']['Template']['aliases'],
>(aliases: E) {
  aliases?.sort()
}
