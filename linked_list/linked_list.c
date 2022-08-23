#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*
 * The data structure that each element of the linked list consists of.
 */
typedef struct linked_list_node {
  void *data;
  struct linked_list_node *child;
} linked_list_node;

/*
 * The wrapper that makes up a whole linked list.
 */
typedef struct linked_list {
  linked_list_node *head;
  linked_list_node *tail;
} linked_list;

/*
 * Add an element to the end of a linked list.
 */
void linked_list_append(linked_list *l, void *data) {
  linked_list_node *n = malloc(sizeof(linked_list_node));
  n->child = NULL;
  n->data = data;
  l->tail->child = n;
  l->tail = n;
}

/*
 * Print out the contents of a linked list.
 */
void _linked_list_print(linked_list_node *n, char *fmt) {
  if (!n) {
    return;
  }
  int i;
  memcpy(&i, n->data, sizeof(int));
  printf(fmt, i);
  _linked_list_print(n->child, fmt);
}

void linked_list_print(linked_list *l, char *fmt) {
  _linked_list_print(l->head->child, fmt);
}

/*
 * Free every element in a linked list.
 */
void free_all(linked_list_node *n) {
  if (!n) {
    return;
  }
  free_all(n->child);
  free(n->data);
  free(n);
}

/*
 * Initialize a linked list struct.
 */
linked_list *linked_list_init() {
  linked_list_node *n = malloc(sizeof(linked_list_node));
  n->child = NULL;
  n->data = NULL;
  linked_list *l = malloc(sizeof(linked_list));
  l->head = n;
  l->tail = n;

  return l;
}

int main() {
  linked_list *l = linked_list_init();

  /*
   * Add ten elements to the linked list.
   */
  for (int i = 0; i < 10; i++) {
    void *d = malloc(sizeof(int));
    memcpy(d, &i, sizeof(int));
    linked_list_append(l, d);
  }

  /*
   * Print out the linked list.
   */
  linked_list_print(l, "%d\n");

  /*
   * Cleanup.
   */
  free_all(l->head);
  free(l);
}
