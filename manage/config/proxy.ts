const proxy = {
    '/api': {
      target:'localhost:8080',
      changeOrigin:true,
      rewrite:(path:string)=>path.replace(/^\/api/,'')
    }
    
  }
export default proxy