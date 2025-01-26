<script>
	import Executables from './lib/pages/Executables.svelte';
	import ExeDetails from './lib/pages/ExecutableDetails.svelte';
	import { Router, Link, Route } from "svelte-routing";
	import LoadingPage from './lib/pages/LoadingPage.svelte';
	import ApplicationInfo from './lib/pages/ApplicationInfo.svelte';

	import '@fontsource/lato';
	import '@fontsource/fira-code';

	export let url = "";
</script>

<Router {url}>
	<div>
		<Route path="/exe-list" component={Executables} />
		<Route path="/exe-details" component={ExeDetails} />
		<Route path="/exe-info"><ApplicationInfo /></Route>
		<Route path="/"><LoadingPage /></Route>
	</div>
</Router>
