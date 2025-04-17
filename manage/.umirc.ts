import { defineConfig } from '@umijs/max';
import routes from './config/route';
import proxy from './config/proxy';
export default defineConfig({
  antd: {},
  access: {},
  model: {},
  initialState: {},
  request: {},
  layout: {
    title: 'blackPig',
  },
  routes: routes,
  npmClient: 'pnpm',
  proxy:proxy
});

