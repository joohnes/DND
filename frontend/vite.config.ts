import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const isDev = false

export default defineConfig({
	build: {
		outDir: 'build',
	},
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
