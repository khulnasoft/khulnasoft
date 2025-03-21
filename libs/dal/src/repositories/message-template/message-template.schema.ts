import mongoose, { Schema } from 'mongoose';
import { ActorTypeEnum } from '@khulnasoft/shared';

import { schemaOptions } from '../schema-default.options';
import { MessageTemplateDBModel } from './message-template.entity';

const mongooseDelete = require('mongoose-delete');

const messageTemplateSchema = new Schema<MessageTemplateDBModel>(
  {
    type: {
      type: Schema.Types.String,
    },
    active: {
      type: Schema.Types.Boolean,
      default: true,
    },
    name: Schema.Types.String,
    stepId: Schema.Types.String,
    subject: Schema.Types.String,
    variables: [
      {
        name: Schema.Types.String,
        type: {
          type: Schema.Types.String,
        },
        required: {
          type: Schema.Types.Boolean,
          default: false,
        },
        defaultValue: Schema.Types.Mixed,
      },
    ],
    content: Schema.Types.Mixed,
    contentType: Schema.Types.String,
    title: Schema.Types.String,
    cta: {
      type: {
        type: Schema.Types.String,
      },
      data: Schema.Types.Mixed,
      action: Schema.Types.Mixed,
    },
    preheader: Schema.Types.String,
    senderName: Schema.Types.String,
    _environmentId: {
      type: Schema.Types.ObjectId,
      ref: 'Environment',
    },
    _organizationId: {
      type: Schema.Types.ObjectId,
      ref: 'Organization',
    },
    _creatorId: {
      type: Schema.Types.ObjectId,
      ref: 'User',
    },
    _feedId: {
      type: Schema.Types.ObjectId,
      ref: 'Feed',
    },
    _parentId: {
      type: Schema.Types.ObjectId,
      ref: 'NotificationTemplate',
    },
    _layoutId: {
      type: Schema.Types.ObjectId,
      ref: 'Layout',
      /*
       * This will make it retro-compatible and will allow
       * that if no layout assigned to not break.
       */
      default: null,
    },
    actor: {
      type: {
        type: Schema.Types.String,
        enum: ActorTypeEnum,
      },
      data: Schema.Types.Mixed,
    },
    controls: { schema: Schema.Types.Mixed, uiSchema: Schema.Types.Mixed },
    output: { schema: Schema.Types.Mixed },
    code: Schema.Types.String,
  },
  schemaOptions
);

messageTemplateSchema.index({
  _organizationId: 1,
  'triggers.identifier': 1,
});

messageTemplateSchema.index({
  _parentId: 1,
});

messageTemplateSchema.plugin(mongooseDelete, { deletedAt: true, deletedBy: true, overrideMethods: 'all' });

export const MessageTemplate =
  (mongoose.models.MessageTemplate as mongoose.Model<MessageTemplateDBModel>) ||
  mongoose.model<MessageTemplateDBModel>('MessageTemplate', messageTemplateSchema);
