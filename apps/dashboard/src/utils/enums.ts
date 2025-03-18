export enum StepTypeEnum {
  IN_APP = 'in_app',
  EMAIL = 'email',
  SMS = 'sms',
  CHAT = 'chat',
  PUSH = 'push',
  DIGEST = 'digest',
  TRIGGER = 'trigger',
  DELAY = 'delay',
  CUSTOM = 'custom',
}

export enum WorkflowTypeEnum {
  REGULAR = 'REGULAR',
  /** @deprecated Use BRIDGE instead */
  ECHO = 'ECHO',
  BRIDGE = 'BRIDGE',
}

export enum WorkflowOriginEnum {
  KHULNASOFT_CLOUD = 'khulnasoft-cloud',
  KHULNASOFT_CLOUD_V1 = 'khulnasoft-cloud-v1',
  EXTERNAL = 'external',
}

export enum WorkflowStatusEnum {
  ACTIVE = 'ACTIVE',
  INACTIVE = 'INACTIVE',
  ERROR = 'ERROR',
}
