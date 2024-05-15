import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const isDev = true

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
		  '/api': {
			target: isDev ? "http://localhost:8000" : 'http://backend:8000',
			changeOrigin: true,
			secure: !isDev,
			rewrite: (path) => path.replace(/^\/api/, ''),
		  },
		},
	  },
});
