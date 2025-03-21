import path from 'path';
import { test } from './utils/baseTest';
import { assertPageShowsMessage, initializeSession, setFeatureFlag } from './utils/browser';
import { TranslationsPage } from './page-models/translationsPage';

test.describe('Translations', () => {
  test.skip(process.env.KHULNASOFT_ENTERPRISE !== 'true', 'Skipping tests for non enterprise variant...');

  test.beforeEach(async ({ page }) => {
    await initializeSession(page, { noTemplates: true });
  });

  test.skip('Can create translation group', async ({ page }) => {
    const groupName = 'My Test Group';
    const translationsPage = await TranslationsPage.goTo(page);
    await translationsPage.assertTitleEquals('Translations');
    await translationsPage.createGroup(groupName);
    await translationsPage.assertHasHeading(groupName);
  });

  test.skip('Can delete a translation group', async ({ page }) => {
    const translationsPage = await TranslationsPage.goTo(page);
    const { identifier } = await translationsPage.createGroup();
    await translationsPage.navigateToGroup(identifier);
    await page.getByTestId('delete-group-btn').click();
    await page.getByTestId('delete-group-submit-btn').click();
    await translationsPage.assertGroupDoesNotExist(identifier);
  });

  test.skip('Can upload translation file', async ({ page }) => {
    const translationsPage = await TranslationsPage.goTo(page);
    const { identifier } = await translationsPage.createGroup();
    await translationsPage.navigateToGroup(identifier);
    await translationsPage.uploadTranslationFile(path.join(__dirname, 'fixtures/translation.json'));
    await assertPageShowsMessage(page, 'translation.json');
  });
});
