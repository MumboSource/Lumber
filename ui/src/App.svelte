<script>
	import Executables from './lib/pages/Executables.svelte';
	import ExeDetails from './lib/pages/ExecutableDetails.svelte';
	import { Router, Link, Route, navigate } from "svelte-routing";
	import LoadingPage from './lib/pages/LoadingPage.svelte';
	import ApplicationInfo from './lib/pages/ApplicationInfo.svelte';
    import SystemInfo from './lib/pages/SystemInfo.svelte';
    import FileEvents from './lib/pages/FileEvents.svelte';

	import '@fontsource/lato';
	import '@fontsource/fira-code';
    import ExecutableDetails from './lib/pages/ExecutableDetails.svelte';
    import FlowChart from './lib/pages/FlowChart.svelte';

	export let url = "";

	let bundles;
	let bundleKeys;

	window.electron.onReceivedBundle((data) => {
		console.log("Rec bundle")

		bundles = data.Apps

		console.log(typeof data.Apps)
		bundleKeys = Object.keys(data.Apps)

        navigate("/exe-list")
    })

</script>

<Router {url}>
	<div>
		<Route path="/exe-list">
			<Executables bundles={bundles} bundleKeys={bundleKeys}/>
		</Route>
		<Route path="/exe-details">
			<ExecutableDetails />
		</Route>
		<Route path="/exe-info">
			<ApplicationInfo />
		</Route>
		<Route path="/sys-info">
			<SystemInfo />
		</Route>
		<Route path="/flow-chart">
			<FlowChart />
		</Route>
		<Route path="/file-events">
			<FileEvents />
		</Route>

		<Route path="/"><LoadingPage /></Route>
	</div>
</Router>
