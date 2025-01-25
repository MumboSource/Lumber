<script>
	import Executables from './lib/executables.svelte';
	import { Router, Link, Route } from "svelte-routing";
	import LoadingPage from './lib/loading-page.svelte';

	export let url = "";
</script>

<Router {url}>
	<div class="w-full h-full">
		<Route path="/exe-list" component={Executables} />
		<Route path="/"><LoadingPage /></Route>
	</div>
</Router>
