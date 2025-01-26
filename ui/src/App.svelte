<script>
	import Executables from './lib/executables.svelte';
	import ExeDetails from './lib/exe-details.svelte';
	import { Router, Link, Route } from "svelte-routing";
	import LoadingPage from './lib/loading-page.svelte';
	import ApplicationInfo from './lib/application-info.svelte';

	import '@fontsource/lato';

	export let url = "";
</script>

<Router {url}>
	<div>
		<Route path="/exe-list" component={Executables} />
		<Route path="/exe-details" component={ExeDetails} />
		<!-- <Route path="/"><LoadingPage /></Route> -->
		<Route path="/"><ApplicationInfo /></Route>
	</div>
</Router>
