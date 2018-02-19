package v1alpha2

import (
	v1alpha2 "github.com/spotahome/redis-operator/api/redisfailover/v1alpha2"
	scheme "github.com/spotahome/redis-operator/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisFailoversGetter has a method to return a RedisFailoverInterface.
// A group's client should implement this interface.
type RedisFailoversGetter interface {
	RedisFailovers(namespace string) RedisFailoverInterface
}

// RedisFailoverInterface has methods to work with RedisFailover resources.
type RedisFailoverInterface interface {
	Create(*v1alpha2.RedisFailover) (*v1alpha2.RedisFailover, error)
	Update(*v1alpha2.RedisFailover) (*v1alpha2.RedisFailover, error)
	UpdateStatus(*v1alpha2.RedisFailover) (*v1alpha2.RedisFailover, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.RedisFailover, error)
	List(opts v1.ListOptions) (*v1alpha2.RedisFailoverList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.RedisFailover, err error)
	RedisFailoverExpansion
}

// redisFailovers implements RedisFailoverInterface
type redisFailovers struct {
	client rest.Interface
	ns     string
}

// newRedisFailovers returns a RedisFailovers
func newRedisFailovers(c *StorageV1alpha2Client, namespace string) *redisFailovers {
	return &redisFailovers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisFailover, and returns the corresponding redisFailover object, and an error if there is any.
func (c *redisFailovers) Get(name string, options v1.GetOptions) (result *v1alpha2.RedisFailover, err error) {
	result = &v1alpha2.RedisFailover{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisFailovers that match those selectors.
func (c *redisFailovers) List(opts v1.ListOptions) (result *v1alpha2.RedisFailoverList, err error) {
	result = &v1alpha2.RedisFailoverList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisFailovers.
func (c *redisFailovers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a redisFailover and creates it.  Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *redisFailovers) Create(redisFailover *v1alpha2.RedisFailover) (result *v1alpha2.RedisFailover, err error) {
	result = &v1alpha2.RedisFailover{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisfailovers").
		Body(redisFailover).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisFailover and updates it. Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *redisFailovers) Update(redisFailover *v1alpha2.RedisFailover) (result *v1alpha2.RedisFailover, err error) {
	result = &v1alpha2.RedisFailover{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(redisFailover.Name).
		Body(redisFailover).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *redisFailovers) UpdateStatus(redisFailover *v1alpha2.RedisFailover) (result *v1alpha2.RedisFailover, err error) {
	result = &v1alpha2.RedisFailover{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(redisFailover.Name).
		SubResource("status").
		Body(redisFailover).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisFailover and deletes it. Returns an error if one occurs.
func (c *redisFailovers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisfailovers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisFailovers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisfailovers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisFailover.
func (c *redisFailovers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.RedisFailover, err error) {
	result = &v1alpha2.RedisFailover{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisfailovers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
