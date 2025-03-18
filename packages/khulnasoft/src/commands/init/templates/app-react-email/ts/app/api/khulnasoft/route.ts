import { serve } from "@khulnasoft/framework/next";
import { welcomeOnboardingEmail } from "../../khulnasoft/workflows";

// the workflows collection can hold as many workflow definitions as you need
export const { GET, POST, OPTIONS } = serve({
  workflows: [welcomeOnboardingEmail],
});
