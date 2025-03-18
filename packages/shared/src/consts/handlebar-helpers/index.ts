export * from './handlebarHelpers';
export * from './getTemplateVariables';

export const khulnasoftReservedVariableNames = ['body'];

export function isReservedVariableName(variableName: string) {
  return khulnasoftReservedVariableNames.includes(variableName);
}
