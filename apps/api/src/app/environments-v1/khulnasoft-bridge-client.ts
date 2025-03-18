import { Inject, Injectable, Scope } from '@nestjs/common';
import type { Request, Response } from 'express';
import { PostActionEnum, type Workflow } from '@khulnasoft/framework/internal';
import { Client, KhulnasoftHandler, KhulnasoftRequestHandler } from '@khulnasoft/framework/nest';
import { GetDecryptedSecretKey, GetDecryptedSecretKeyCommand } from '@khulnasoft/application-generic';
import { ConstructFrameworkWorkflow, ConstructFrameworkWorkflowCommand } from './usecases/construct-framework-workflow';

/*
 * A custom framework name is specified for the Khulnasoft-managed Bridge endpoint
 * to provide a clear distinction between Khulnasoft-managed and self-managed Bridge endpoints.
 */
export const frameworkName = 'khulnasoft-nest';

/**
 * This class overrides the default NestJS Khulnasoft Bridge Client to allow for dynamic construction of
 * workflows to serve on the Khulnasoft Bridge.
 */
@Injectable({ scope: Scope.REQUEST })
export class KhulnasoftBridgeClient {
  public khulnasoftRequestHandler: KhulnasoftRequestHandler | null = null;

  constructor(
    @Inject(KhulnasoftHandler) private khulnasoftHandler: KhulnasoftHandler,
    private constructFrameworkWorkflow: ConstructFrameworkWorkflow,
    private getDecryptedSecretKey: GetDecryptedSecretKey
  ) {}

  public async handleRequest(req: Request, res: Response) {
    const secretKey = await this.getDecryptedSecretKey.execute(
      GetDecryptedSecretKeyCommand.create({
        environmentId: req.params.environmentId,
      })
    );

    const workflows: Workflow[] = [];

    /*
     * Only construct a workflow when dealing with a POST request to the Khulnasoft-managed Bridge endpoint.
     * Non-POST requests don't have a `workflowId` query parameter, so we can't construct a workflow.
     * Those non-POST requests are handled for the purpose of returning a successful health-check.
     */
    if (Object.values(PostActionEnum).includes(req.query.action as PostActionEnum)) {
      const programmaticallyConstructedWorkflow = await this.constructFrameworkWorkflow.execute(
        ConstructFrameworkWorkflowCommand.create({
          environmentId: req.params.environmentId,
          workflowId: req.query.workflowId as string,
          controlValues: req.body.controls,
          action: req.query.action as PostActionEnum,
        })
      );

      workflows.push(programmaticallyConstructedWorkflow);
    }

    this.khulnasoftRequestHandler = new KhulnasoftRequestHandler({
      frameworkName,
      workflows,
      client: new Client({ secretKey, strictAuthentication: true }),
      handler: this.khulnasoftHandler.handler,
    });

    await this.khulnasoftRequestHandler.createHandler()(req, res);
  }
}
