# pizza-controller-me

1. apply the manifest

```
git clone https://github.com/jansure/pizza-controller-me
cd pizza-controller-me

# using `kapp` - recommended :)
#
kapp deploy -a pizza-controller -f ./config/release.yaml


# OR .. plain old kubectl
#
kubectl apply -f ./config/release.yaml
```
