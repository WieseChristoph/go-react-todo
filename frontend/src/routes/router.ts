import { RootRoute, Route, Router } from "@tanstack/react-router";

import Root from "./Root";
import Home from "./Home";
import Todos from "./Todos";
import Profile from "./Profile";

const rootRoute = new RootRoute({
  component: Root,
});

const routeTree = rootRoute.addChildren([
  new Route({
    getParentRoute: () => rootRoute,
    path: "/",
    component: Home,
  }),
  new Route({
    getParentRoute: () => rootRoute,
    path: "/todos",
    component: Todos,
  }),
  new Route({
    getParentRoute: () => rootRoute,
    path: "/profile",
    component: Profile,
  }),
]);

export default new Router({ routeTree });
