import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { KhulnasoftModule } from '@khulnasoft/framework/nest';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { NotificationService } from './notification.service';
import { UserService } from './user.service';

@Module({
  imports: [
    /*
     * IMPORTANT: ConfigModule must be imported before KhulnasoftModule to ensure
     * environment variables are loaded before the KhulnasoftModule is initialized.
     *
     * This ensures that KHULNASOFT_SECRET_KEY is available when the KhulnasoftModule is initialized.
     */
    ConfigModule.forRoot({
      envFilePath: '.env',
    }),
    KhulnasoftModule.registerAsync({
      imports: [AppModule],
      useFactory: (notificationService: NotificationService) => ({
        apiPath: '/api/khulnasoft',
        workflows: [notificationService.welcomeWorkflow()],
      }),
      inject: [NotificationService],
    }),
  ],
  controllers: [AppController],
  providers: [AppService, NotificationService, UserService],
  exports: [NotificationService],
})
export class AppModule {}
