import React from 'react'
// import { AutoForm } from 'uniforms'
// @ts-ignore
import MarkdownExample from './../../lib/MarkdownSample.md'
import ReactMarkdown from 'react-markdown'
const gfm = require('remark-gfm')

import Typography from '.'
// @ts-ignore
import { Space } from '../../index'

const { Title, Text, Link } = Typography

export default {
  title: 'General/Typography',
  component: Typography,
}

export const article = () => (
  <Typography tag="article">
    <ReactMarkdown plugins={[gfm]} source={MarkdownExample} />
  </Typography>
)

export const Titles = () => (
  <React.Fragment>
    <Title level={1}>Hello world</Title>
    <Title level={2}>Hello world</Title>
    <Title level={3}>Hello world</Title>
    <Title level={4}>Hello world</Title>
    <Title level={5}>Hello world</Title>
  </React.Fragment>
)

export const Texts = () => (
  <>
    <Text>Khulnasoft UI (default)</Text>
    <br />
    <Text type="secondary">Khulnasoft UI (secondary)</Text>
    <br />
    <Text type="success">Khulnasoft UI (success)</Text>
    <br />
    <Text type="warning">Khulnasoft UI (warning)</Text>
    <br />
    <Text type="danger">Khulnasoft UI (danger)</Text>
    <br />
    <Text disabled>Khulnasoft UI (disabled)</Text>
    <br />
    <Text mark>Khulnasoft UI (mark)</Text>
    <br />
    <Text code>Khulnasoft UI (code)</Text>
    <br />
    <Text keyboard>Khulnasoft UI (keyboard)</Text>
    <br />
    <Text underline>Khulnasoft UI (underline)</Text>
    <br />
    <Text strikethrough>Khulnasoft UI (strikethrough)</Text>
    <br />
    <Text strong>Khulnasoft UI (strong)</Text>
    <br />
    <Link href="https://khulnasoft.io" target="_blank">
      Khulnasoft (Link)
    </Link>
  </>
)
