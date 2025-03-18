import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import request from 'supertest';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { AppModule } from './app.module';

describe('AppController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    process.env.KHULNASOFT_SECRET_KEY = 'test';
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  afterEach(async () => {
    await app.close();
  });

  it('/ (GET)', () => {
    return request(app.getHttpServer()).get('/').expect(200).expect('Hello World!');
  });

  it('/api/khulnasoft (GET)', async () => {
    const response = await request(app.getHttpServer()).get('/api/khulnasoft?action=health-check').expect(200);

    expect(response.body).toEqual(
      expect.objectContaining({
        status: 'ok',
        sdkVersion: expect.any(String),
        frameworkVersion: expect.any(String),
        discovered: { workflows: 1, steps: 1 },
      })
    );
  });
});
