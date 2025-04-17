
const route = [
    {
      path: '/',
      redirect: '/home',
    },
    {
      name: '首页',
      path: '/home',
      component: './Home',
    },
    {
        name: '预约列表',
        path: '/bookList',
        component: './BookList'
    },
    {
        name: '预约设置',
        path: '/book',
        component: './Book'
    },

  ]
export default route