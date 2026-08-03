# Block 5 - Scaling Types (14:58-17:38)

| Type | What it does | Analogy |
|---|---|---|
| Horizontal | Add more machines to the cluster | More trucks |
| Vertical | Enlarge the machine (more RAM, CPU, disk) | A bigger truck |

In practice, both are combined. Since every machine costs money, the goal is to use the minimum necessary.

## Autoscaling

The load balancer detects traffic spikes and provisions instances automatically (for example, a minimum of 2 and a maximum of 10). This is done by Kubernetes and cloud providers.

## Serverless

You do not worry about provisioning the server; you hand over a Docker container and the provider takes care of the rest.
