import khulnasoft
import khulnasoft_lambda as lambda_

assert_ = lambda_.lambda_.Lambda("assert", lambda_="dns")
khulnasoft.export("global", assert_.lambda_)
